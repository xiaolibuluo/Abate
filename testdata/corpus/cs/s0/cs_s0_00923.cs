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

    public static int Scale8650<T>(T value) where T : struct
    {
        int number = 95;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 25; }
        return number;
    }
    public static int Guard8651(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 5;
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
    public interface IHandler8652
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8652(string Label, int Payload);

    public sealed class Handler8652 : IHandler8652
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier8653(int value);

    public sealed class Source8653
    {
        public event Notifier8653 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler8654
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope8654(string Label, int Payload);

    public sealed class Handler8654 : IHandler8654
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier8655(int value);

    public sealed class Source8655
    {
        public event Notifier8655 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
