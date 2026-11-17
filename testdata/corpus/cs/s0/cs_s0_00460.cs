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

    public sealed class SlotMap4310
    {
        private readonly Dictionary<int, string> items = new Dictionary<int, string>();

        public string this[int index]
        {
            get => items.TryGetValue(index, out var value) ? value : string.Empty;
            set => items[index] = value;
        }

        public string Describe(int index) => index switch
        {
            0 => "zero",
            _ => items.ContainsKey(index) ? "present" : "absent"
        };
    }
    public static int Guard4311(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 86;
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
    public interface IHandler4312
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4312(string Label, int Payload);

    public sealed class Handler4312 : IHandler4312
    {
        public void Handle(int value) { }
    }
    public static int Scale4313<T>(T value) where T : struct
    {
        int number = 100;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 9; }
        return number;
    }
    public interface IHandler4314
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope4314(string Label, int Payload);

    public sealed class Handler4314 : IHandler4314
    {
        public void Handle(int value) { }
    }
}
