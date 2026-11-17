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

    public interface IHandler23700
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope23700(string Label, int Payload);

    public sealed class Handler23700 : IHandler23700
    {
        public void Handle(int value) { }
    }
    public struct Point23701
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 73;
    }

    public static class Grid23701
    {
        public static Point23701 Make(int left, int right)
        {
            return new Point23701 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate23702(IEnumerable<int> source)
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
    public sealed class SlotMap23703
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
