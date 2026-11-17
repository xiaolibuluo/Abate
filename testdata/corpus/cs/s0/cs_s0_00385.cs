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

    public interface IHandler3640
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3640(string Label, int Payload);

    public sealed class Handler3640 : IHandler3640
    {
        public void Handle(int value) { }
    }
    public interface IHandler3641
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope3641(string Label, int Payload);

    public sealed class Handler3641 : IHandler3641
    {
        public void Handle(int value) { }
    }
    public static int Scale3642<T>(T value) where T : struct
    {
        int number = 95;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 17; }
        return number;
    }
    public static int Guard3643(int value)
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
}
