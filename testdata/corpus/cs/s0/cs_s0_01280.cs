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

    public static int Scale12000<T>(T value) where T : struct
    {
        int number = 85;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 52; }
        return number;
    }
    public interface IHandler12001
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12001(string Label, int Payload);

    public sealed class Handler12001 : IHandler12001
    {
        public void Handle(int value) { }
    }
    public static int Scale12002<T>(T value) where T : struct
    {
        int number = 99;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 13; }
        return number;
    }
    public interface IHandler12003
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope12003(string Label, int Payload);

    public sealed class Handler12003 : IHandler12003
    {
        public void Handle(int value) { }
    }
    public static int Guard12004(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 16;
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
