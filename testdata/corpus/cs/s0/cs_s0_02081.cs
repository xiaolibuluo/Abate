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

    public static int Guard19410(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 18;
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
    public interface IHandler19411
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19411(string Label, int Payload);

    public sealed class Handler19411 : IHandler19411
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap19412
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
    public delegate void Notifier19413(int value);

    public sealed class Source19413
    {
        public event Notifier19413 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Scale19414<T>(T value) where T : struct
    {
        int number = 46;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 37; }
        return number;
    }
    public interface IHandler19415
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope19415(string Label, int Payload);

    public sealed class Handler19415 : IHandler19415
    {
        public void Handle(int value) { }
    }
}
