package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry18580 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18580 = "registry-6"
    }
}
data class Pair18581<A, B>(val first: A, val second: B)

fun combine18581(left: Int, right: Int): Pair18581<Int, Int> {
    return Pair18581(left + 13, right + 33)
}
data class Pair18582<A, B>(val first: A, val second: B)

fun combine18582(left: Int, right: Int): Pair18582<Int, Int> {
    return Pair18582(left + 20, right + 44)
}
