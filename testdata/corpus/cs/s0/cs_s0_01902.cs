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

    public interface IHandler17750
    {
        void Handle(int value);
    }

    [Serializable]
    public sealed record Envelope17750(string Label, int Payload);

    public sealed class Handler17750 : IHandler17750
    {
        public void Handle(int value) { }
    }
    public struct Point17751
    {
        public int X { get; set; }
        public int Y { get; set; }
        public int Total => X + Y + 98;
    }

    public static class Grid17751
    {
        public static Point17751 Make(int left, int right)
        {
            return new Point17751 { X = left, Y = right };
        }
    }
    public async Task<int> Aggregate17752(IEnumerable<int> source)
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
