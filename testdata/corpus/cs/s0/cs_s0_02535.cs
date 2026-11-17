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

    public static int Guard23680(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 40;
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
    public static int Guard23681(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 47;
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
    public delegate void Notifier23682(int value);

    public sealed class Source23682
    {
        public event Notifier23682 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale23683<T>(T value) where T : struct
    {
        int number = 61;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 26; }
        return number;
    }
    public static int Scale23684<T>(T value) where T : struct
    {
        int number = 68;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 37; }
        return number;
    }
}
