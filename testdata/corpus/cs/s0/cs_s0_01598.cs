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

    public delegate void Notifier14940(int value);

    public sealed class Source14940
    {
        public event Notifier14940 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale14941<T>(T value) where T : struct
    {
        int number = 34;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 47; }
        return number;
    }
    public static int Scale14942<T>(T value) where T : struct
    {
        int number = 41;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 58; }
        return number;
    }
    public static int Guard14943(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 48;
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
    public static int Guard14944(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 55;
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
}
