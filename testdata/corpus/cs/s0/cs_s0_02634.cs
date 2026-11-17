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

    public interface IHandler24640
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24640(string Label, int Payload);

    public sealed class Handler24640 : IHandler24640
    {
        public void Handle(int value) { }
    }
    public interface IHandler24641
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24641(string Label, int Payload);

    public sealed class Handler24641 : IHandler24641
    {
        public void Handle(int value) { }
    }
    public static int Scale24642<T>(T value) where T : struct
    {
        int number = 41;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 16; }
        return number;
    }
    public static int Guard24643(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 48;
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
