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

    public delegate void Notifier10050(int value);

    public sealed class Source10050
    {
        public event Notifier10050 Changed;
        public void Raise(int value) { Changed?.Invoke(value); }
    }
    public struct Point10051
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 79;
    }

    public static class Grid10051
    {
        public static Point10051 Make(int left, int right)
        {
            return new Point10051 { X = left, Y = right };
        }
    }
    public sealed class SlotMap10052
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
    public async Task<int> Aggregate10053(IEnumerable<int> source)
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
}
