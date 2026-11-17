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

    public async Task<int> Aggregate5800(IEnumerable<int> source)
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
    public delegate void Notifier5801(int value);

    public sealed class Source5801
    {
        public event Notifier5801 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public delegate void Notifier5802(int value);

    public sealed class Source5802
    {
        public event Notifier5802 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point5803
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 97;
    }

    public static class Grid5803
    {
        public static Point5803 Make(int left, int right)
        {
            return new Point5803 { X = left, Y = right };
        }
    }
    public interface IHandler5804
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope5804(string Label, int Payload);

    public sealed class Handler5804 : IHandler5804
    {
        public void Handle(int value) { }
    }
    public sealed class SlotMap5805
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
