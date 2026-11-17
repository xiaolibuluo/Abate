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

    public static int Guard17940(int value)
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
    public interface IHandler17941
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17941(string Label, int Payload);

    public sealed class Handler17941 : IHandler17941
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier17942(int value);

    public sealed class Source17942
    {
        public event Notifier17942 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale17943<T>(T value) where T : struct
    {
        int number = 68;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 34; }
        return number;
    }
    public interface IHandler17944
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17944(string Label, int Payload);

    public sealed class Handler17944 : IHandler17944
    {
        public void Handle(int value) { }
    }
    public static int Scale17945<T>(T value) where T : struct
    {
        int number = 82;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 56; }
        return number;
    }
}
