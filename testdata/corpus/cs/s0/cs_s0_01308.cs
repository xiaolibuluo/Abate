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

    public static int Scale12260<T>(T value) where T : struct
    {
        int number = 35;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 51; }
        return number;
    }
    public static int Scale12261<T>(T value) where T : struct
    {
        int number = 42;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 62; }
        return number;
    }
}
