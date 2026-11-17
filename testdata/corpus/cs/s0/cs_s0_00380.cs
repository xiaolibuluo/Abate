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

    public static int Scale3590<T>(T value) where T : struct
    {
        int number = 16;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 21; }
        return number;
    }
    public static int Scale3591<T>(T value) where T : struct
    {
        int number = 23;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 32; }
        return number;
    }
    public static int Scale3592<T>(T value) where T : struct
    {
        int number = 30;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 43; }
        return number;
    }
}
