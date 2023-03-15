import { Box, Button, Toolbar, Typography } from "@mui/material"
import AppBar from "@mui/material/AppBar"
import { Archive, Home } from "@mui/icons-material";
import AttachFileIcon from '@mui/icons-material/AttachFile';
import { useNavigate } from "react-router-dom";

const Header = () => {

    const navigate = useNavigate();

    return (
        <div>
            <Box>
                <AppBar position="fixed">
                    <Toolbar>
                        <AttachFileIcon sx={{ display: { xs: 'none', md: 'flex' }, mr: 1 }} />
                        <Typography variant="h6" component="div" sx={{ flexGrow: 0.02 }}>
                            Bind
                        </Typography>
                        <Box sx={{ display: { xs: 'none', md: 'flex' } }}>
                            <Button startIcon={<Home />} sx={{ color: '#fff' }} onClick={() => navigate("/")}>
                                Home
                            </Button>
                            <Button startIcon={<Archive />} sx={{ color: '#fff' }} onClick={() => navigate("/archive")}>
                                Archive
                            </Button>
                        </Box>
                    </Toolbar>
                </AppBar>
            </Box>
        </div>
    )
}

export default Header