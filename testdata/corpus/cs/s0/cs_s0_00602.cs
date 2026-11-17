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

    public delegate void Notifier5630(int value);

    public sealed class Source5630
    {
        public event Notifier5630 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler5631
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5631(string Label, int Payload);

    public sealed class Handler5631 : IHandler5631
    {
        public void Handle(int value) { }
    }
    public static int Guard5632(int value)
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
    public interface IHandler5633
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5633(string Label, int Payload);

    public sealed class Handler5633 : IHandler5633
    {
        public void Handle(int value) { }
    }
}
