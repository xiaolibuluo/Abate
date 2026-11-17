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

    public static int Guard18910(int value)
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
    public interface IHandler18911
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18911(string Label, int Payload);

    public sealed class Handler18911 : IHandler18911
    {
        public void Handle(int value) { }
    }
    public static int Guard18912(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 61;
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
    public delegate void Notifier18913(int value);

    public sealed class Source18913
    {
        public event Notifier18913 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler18914
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18914(string Label, int Payload);

    public sealed class Handler18914 : IHandler18914
    {
        public void Handle(int value) { }
    }
    public interface IHandler18915
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope18915(string Label, int Payload);

    public sealed class Handler18915 : IHandler18915
    {
        public void Handle(int value) { }
    }
}
