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

    public interface IHandler25570
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope25570(string Label, int Payload);

    public sealed class Handler25570 : IHandler25570
    {
        public void Handle(int value) { }
    }
    public async Task<int> Aggregate25571(IEnumerable<int> source)
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
    public struct Point25572
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 86;
    }

    public static class Grid25572
    {
        public static Point25572 Make(int left, int right)
        {
            return new Point25572 { X = left, Y = right };
        }
    }
}
