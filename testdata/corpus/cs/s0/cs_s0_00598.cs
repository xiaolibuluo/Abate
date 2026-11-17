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

    public static int Scale5590<T>(T value) where T : struct
    {
        int number = 94;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 18; }
        return number;
    }
    public static int Guard5591(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 101;
        }
        catch (ArgumentException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("guard");
        }
    }
    public static int Scale5592<T>(T value) where T : struct
    {
        int number = 11;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 40; }
        return number;
    }
    public static int Scale5593<T>(T value) where T : struct
    {
        int number = 18;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 51; }
        return number;
    }
}
