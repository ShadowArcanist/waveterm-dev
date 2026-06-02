// Copyright 2026, Command Line Inc.
// SPDX-License-Identifier: Apache-2.0

import { ClientModel } from "@/app/store/client-model";
import { globalStore } from "@/app/store/jotaiStore";
import * as services from "@/app/store/services";
import { atoms } from "@/store/global";
import { modalsModel } from "@/store/modalmodel";
import { fireAndForget } from "@/util/util";
import * as jotai from "jotai";
import { useEffect } from "react";
import { getModalComponent } from "./modalregistry";

// Guard so the first-launch bootstrap runs at most once per session. AgreeTos ->
// BootstrapStarterLayout is NOT idempotent (it re-applies the starter layout), and this
// effect can otherwise fire multiple times (React StrictMode, or clientData changing before
// the tosagreed flag round-trips back through WOS), which would duplicate the starter blocks.
let tosAgreeRequested = false;

const ModalsRenderer = () => {
    const clientData = jotai.useAtomValue(ClientModel.getInstance().clientAtom);
    const [modals] = jotai.useAtom(modalsModel.modalsAtom);
    const rtn: React.ReactElement[] = [];
    for (const modal of modals) {
        const ModalComponent = getModalComponent(modal.displayName);
        if (ModalComponent) {
            rtn.push(<ModalComponent key={modal.displayName} {...modal.props} />);
        }
    }
    // Onboarding has been removed. On first launch we silently bootstrap the
    // starter layout (this is what ToS agreement used to do) with no wizard.
    useEffect(() => {
        if (!clientData.tosagreed && !tosAgreeRequested) {
            tosAgreeRequested = true;
            fireAndForget(() => services.ClientService.AgreeTos());
        }
    }, [clientData]);
    useEffect(() => {
        globalStore.set(atoms.modalOpen, rtn.length > 0);
    }, [rtn]);

    return <>{rtn}</>;
};

export { ModalsRenderer };
