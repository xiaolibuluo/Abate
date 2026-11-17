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

    public delegate void Notifier14340(int value);

    public sealed class Source14340
    {
        public event Notifier14340 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler14341
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14341(string Label, int Payload);

    public sealed class Handler14341 : IHandler14341
    {
        public void Handle(int value) { }
    }
    public static int Scale14342<T>(T value) where T : struct
    {
        int number = 37;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 65; }
        return number;
    }
    public static int Guard14343(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 44;
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
