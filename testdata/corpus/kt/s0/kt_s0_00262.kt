package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair2140<A, B>(val first: A, val second: B)

fun combine2140(left: Int, right: Int): Pair2140<Int, Int> {
    return Pair2140(left + 71, right + 43)
}
data class Pair2141<A, B>(val first: A, val second: B)

fun combine2141(left: Int, right: Int): Pair2141<Int, Int> {
    return Pair2141(left + 78, right + 54)
}
object Registry2142 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2142 = "registry-85"
    }
}
object Registry2143 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG2143 = "registry-92"
    }
}
fun <T> List<T>.secondOrNull2144(): T? = if (size > 1) this[1] else null

fun render2144(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
