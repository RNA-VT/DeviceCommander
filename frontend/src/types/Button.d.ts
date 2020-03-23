import { ButtonClassKey } from "@material-ui/core/Button/Button";
import { StandardProps } from "@material-ui/core/index";

declare module "material-ui/Button/Button" {
    export interface Button extends StandardProps<{}, ButtonClassKey> {
        onClick?: () => void;
    }
}