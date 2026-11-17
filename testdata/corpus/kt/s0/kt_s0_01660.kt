package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry13670 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG13670 = "registry-25"
    }
}
data class Pair13671<A, B>(val first: A, val second: B)

fun combine13671(left: Int, right: Int): Pair13671<Int, Int> {
    return Pair13671(left + 32, right + 12)
}
