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

    public interface IHandler14880
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14880(string Label, int Payload);

    public sealed class Handler14880 : IHandler14880
    {
        public void Handle(int value) { }
    }
    public static int Guard14881(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 53;
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
    public interface IHandler14882
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope14882(string Label, int Payload);

    public sealed class Handler14882 : IHandler14882
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier14883(int value);

    public sealed class Source14883
    {
        public event Notifier14883 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
