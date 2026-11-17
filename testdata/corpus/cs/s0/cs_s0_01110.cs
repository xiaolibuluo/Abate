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

    public static int Guard10400(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 42;
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
    public delegate void Notifier10401(int value);

    public sealed class Source10401
    {
        public event Notifier10401 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard10402(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 56;
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
