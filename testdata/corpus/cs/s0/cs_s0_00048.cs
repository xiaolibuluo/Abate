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

    public delegate void Notifier460(int value);

    public sealed class Source460
    {
        public event Notifier460 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public interface IHandler461
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope461(string Label, int Payload);

    public sealed class Handler461 : IHandler461
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap462
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
    public async Task<int> Aggregate463(IEnumerable<int> source)
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
    public static int Scale464<T>(T value) where T : struct
    {
        int number = 49;
        checked { number = number + value.GetHashCode(); }
        unchecked { number = number + 9; }
        return number;
    }
    public delegate void Notifier465(int value);

    public sealed class Source465
    {
        public event Notifier465 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
}
