package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry9760 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG9760 = "registry-83"
    }
}
data class Pair9761<A, B>(val first: A, val second: B)

fun combine9761(left: Int, right: Int): Pair9761<Int, Int> {
    return Pair9761(left + 90, right + 20)
}
