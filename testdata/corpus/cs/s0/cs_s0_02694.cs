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

    public static int Scale25190<T>(T value) where T : struct
    {
        int number = 63;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 13; }
        return number;
    }
    public interface IHandler25191
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25191(string Label, int Payload);

    public sealed class Handler25191 : IHandler25191
    {
        public void Handle(int value) { }
    }
    public static int Guard25192(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 77;
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
    public interface IHandler25193
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25193(string Label, int Payload);

    public sealed class Handler25193 : IHandler25193
    {
        public void Handle(int value) { }
    }
    public interface IHandler25194
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25194(string Label, int Payload);

    public sealed class Handler25194 : IHandler25194
    {
        public void Handle(int value) { }
    }
}
