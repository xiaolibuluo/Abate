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

    public interface IHandler1880
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope1880(string Label, int Payload);

    public sealed class Handler1880 : IHandler1880
    {
        public void Handle(int value) { }
    }
    public static int Guard1881(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 31;
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
    public interface IHandler1882
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope1882(string Label, int Payload);

    public sealed class Handler1882 : IHandler1882
    {
        public void Handle(int value) { }
    }
    public static int Scale1883<T>(T value) where T : struct
    {
        int number = 45;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 16; }
        return number;
    }
    public sealed class SlotMap1884
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
}
