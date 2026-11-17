package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair16100<A, B>(val first: A, val second: B)

fun combine16100(left: Int, right: Int): Pair16100<Int, Int> {
    return Pair16100(left + 80, right + 55)
}
object Registry16101 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16101 = "registry-87"
    }
}
