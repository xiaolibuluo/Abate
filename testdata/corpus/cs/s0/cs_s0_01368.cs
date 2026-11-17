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

    public static int Guard12820(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 84;
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
    public static int Scale12821<T>(T value) where T : struct
    {
        int number = 91;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 27; }
        return number;
    }
    public delegate void Notifier12822(int value);

    public sealed class Source12822
    {
        public event Notifier12822 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
