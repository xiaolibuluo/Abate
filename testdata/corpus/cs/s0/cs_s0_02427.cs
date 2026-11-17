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

    public static int Guard22670(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 85;
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
    public interface IHandler22671
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22671(string Label, int Payload);

    public sealed class Handler22671 : IHandler22671
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier22672(int value);

    public sealed class Source22672
    {
        public event Notifier22672 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale22673<T>(T value) where T : struct
    {
        int number = 9;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 51; }
        return number;
    }
}
