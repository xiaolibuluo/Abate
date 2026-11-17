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

    public static int Scale24230<T>(T value) where T : struct
    {
        int number = 76;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 12; }
        return number;
    }
    public interface IHandler24231
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24231(string Label, int Payload);

    public sealed class Handler24231 : IHandler24231
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier24232(int value);

    public sealed class Source24232
    {
        public event Notifier24232 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard24233(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 97;
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
    public interface IHandler24234
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope24234(string Label, int Payload);

    public sealed class Handler24234 : IHandler24234
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap24235
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
