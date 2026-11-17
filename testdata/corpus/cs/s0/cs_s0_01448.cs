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

    public static int Guard13570(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 89;
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
    public static int Scale13571<T>(T value) where T : struct
    {
        int number = 96;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 64; }
        return number;
    }
    public sealed class SlotMap13572
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
    public interface IHandler13573
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope13573(string Label, int Payload);

    public sealed class Handler13573 : IHandler13573
    {
        public void Handle(int value) { }
    }
    public delegate void Notifier13574(int value);

    public sealed class Source13574
    {
        public event Notifier13574 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public static int Guard13575(int value)
    {
        try
        {
            if (value < 0) { throw new ArgumentException("bad value"); }
            return value + 27;
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
