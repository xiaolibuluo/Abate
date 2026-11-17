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

    public static int Guard19630(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 13;
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
    public delegate void Notifier19631(int value);

    public sealed class Source19631
    {
        public event Notifier19631 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler19632
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19632(string Label, int Payload);

    public sealed class Handler19632 : IHandler19632
    {
        public void Handle(int value) { }
    }
}
