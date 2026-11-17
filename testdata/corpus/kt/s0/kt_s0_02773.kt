package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair22890<A, B>(val first: A, val second: B)

fun combine22890(left: Int, right: Int): Pair22890<Int, Int> {
    return Pair22890(left + 80, right + 50)
}
object Registry22891 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG22891 = "registry-87"
    }
}
