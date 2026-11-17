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

    public static int Scale16640<T>(T value) where T : struct
    {
        int number = 6;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 67; }
        return number;
    }
    public static int Scale16641<T>(T value) where T : struct
    {
        int number = 13;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 17; }
        return number;
    }
}
