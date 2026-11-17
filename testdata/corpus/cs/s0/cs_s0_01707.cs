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

    public static int Guard15950(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 79;
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
    public sealed class SlotMap15951
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
    public delegate void Notifier15952(int value);

    public sealed class Source15952
    {
        public event Notifier15952 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale15953<T>(T value) where T : struct
    {
        int number = 100;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 44; }
        return number;
    }
    public interface IHandler15954
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope15954(string Label, int Payload);

    public sealed class Handler15954 : IHandler15954
    {
        public void Handle(int value) { }
    }
}
