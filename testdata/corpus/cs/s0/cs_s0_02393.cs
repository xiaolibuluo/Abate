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

    public interface IHandler22350
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope22350(string Label, int Payload);

    public sealed class Handler22350 : IHandler22350
    {
        public void Handle(int value) { }
    }
    public static int Scale22351<T>(T value) where T : struct
    {
        int number = 64;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 49; }
        return number;
    }
    public delegate void Notifier22352(int value);

    public sealed class Source22352
    {
        public event Notifier22352 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier22353(int value);

    public sealed class Source22353
    {
        public event Notifier22353 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public async Task<int> Aggregate22354(IEnumerable<int> source)
    {
        try
        {
            using var cursor = source.GetEnumerator();
            return await Task.FromResult(source.Where(item => item > 0).Sum());
        }
        catch (InvalidOperationException)
        {
            return 0;
        }
        finally
        {
            Console.WriteLine("done");
        }
    }
    public static int Scale22355<T>(T value) where T : struct
    {
        int number = 92;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 32; }
        return number;
    }
    public sealed class SlotMap22356
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
