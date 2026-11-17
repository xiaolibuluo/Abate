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

    public static int Guard11610(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 63;
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
    public delegate void Notifier11611(int value);

    public sealed class Source11611
    {
        public event Notifier11611 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier11612(int value);

    public sealed class Source11612
    {
        public event Notifier11612 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale11613<T>(T value) where T : struct
    {
        int number = 84;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 56; }
        return number;
    }
    public sealed class SlotMap11614
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
    public interface IHandler11615
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope11615(string Label, int Payload);

    public sealed class Handler11615 : IHandler11615
    {
        public void Handle(int value) { }
    }
}
