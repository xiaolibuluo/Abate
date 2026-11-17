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

    public interface IHandler17920
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17920(string Label, int Payload);

    public sealed class Handler17920 : IHandler17920
    {
        public void Handle(int value) { }
    }
    public static int Scale17921<T>(T value) where T : struct
    {
        int number = 28;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 59; }
        return number;
    }
    public interface IHandler17922
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17922(string Label, int Payload);

    public sealed class Handler17922 : IHandler17922
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap17923
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
    public static int Guard17924(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 49;
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
