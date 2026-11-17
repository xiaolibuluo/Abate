package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry1770 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1770 = "registry-75"
    }
}
object Registry1771 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1771 = "registry-82"
    }
}
data class Pair1772<A, B>(val first: A, val second: B)

fun combine1772(left: Int, right: Int): Pair1772<Int, Int> {
    return Pair1772(left + 89, right + 50)
}
data class Pair1773<A, B>(val first: A, val second: B)

fun combine1773(left: Int, right: Int): Pair1773<Int, Int> {
    return Pair1773(left + 96, right + 61)
}
object Registry1774 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1774 = "registry-6"
    }
}
object Registry1775 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG1775 = "registry-13"
    }
}
@JvmName("step1776")
fun step1776(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 20
}
fun <T> List<T>.secondOrNull1777(): T? = if (size > 1) this[1] else null

fun render1777(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
