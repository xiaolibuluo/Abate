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

    public static int Scale3150<T>(T value) where T : struct
    {
        int number = 26;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 18; }
        return number;
    }
    public interface IHandler3151
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3151(string Label, int Payload);

    public sealed class Handler3151 : IHandler3151
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier3152(int value);

    public sealed class Source3152
    {
        public event Notifier3152 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard3153(int value)
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
}
