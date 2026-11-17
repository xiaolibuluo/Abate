package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry18550 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG18550 = "registry-64"
    }
}
data class Pair18551<A, B>(val first: A, val second: B)

fun combine18551(left: Int, right: Int): Pair18551<Int, Int> {
    return Pair18551(left + 71, right + 12)
}
