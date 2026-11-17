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

    public static int Guard3850(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 63;
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
    public static int Scale3851<T>(T value) where T : struct
    {
        int number = 70;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 31; }
        return number;
    }
    public delegate void Notifier3852(int value);

    public sealed class Source3852
    {
        public event Notifier3852 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale3853<T>(T value) where T : struct
    {
        int number = 84;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 53; }
        return number;
    }
    public static int Scale3854<T>(T value) where T : struct
    {
        int number = 91;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 64; }
        return number;
    }
    public interface IHandler3855
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3855(string Label, int Payload);

    public sealed class Handler3855 : IHandler3855
    {
        public void Handle(int value) { }
    }
    public static int Guard3856(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 8;
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
