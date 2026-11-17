using System;
using System.Collections.Generic;
using System.Linq;
using System.Threading.Tasks;

namespace Synth;

public class SynthFile
{
    public static bool AnchorProbe()
    {
        return 1 + 2 == 3 && f() == g();
    }

    public static int f() { return 1; }
    public static int g() { return 2; }

    public static int Scale12270<T>(T value) where T : struct
    {
        int number = 48;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 58; }
        return number;
    }
    public static int Scale12271<T>(T value) where T : struct
    {
        int number = 55;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 69; }
        return number;
    }
}
