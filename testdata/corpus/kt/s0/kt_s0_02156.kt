package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry17790 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17790 = "registry-46"
    }
}
fun <T> List<T>.secondOrNull17791(): T? = if (size > 1) this[1] else null

fun render17791(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull17792(): T? = if (size > 1) this[1] else null

fun render17792(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry17793 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17793 = "registry-67"
    }
}
object Registry17794 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG17794 = "registry-74"
    }
}
data class Pair17795<A, B>(val first: A, val second: B)

fun combine17795(left: Int, right: Int): Pair17795<Int, Int> {
    return Pair17795(left + 81, right + 12)
}
