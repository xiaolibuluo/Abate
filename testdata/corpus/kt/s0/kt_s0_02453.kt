package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry20150 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20150 = "registry-10"
    }
}
fun <T> List<T>.secondOrNull20151(): T? = if (size > 1) this[1] else null

fun render20151(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair20152<A, B>(val first: A, val second: B)

fun combine20152(left: Int, right: Int): Pair20152<Int, Int> {
    return Pair20152(left + 24, right + 45)
}
data class Pair20153<A, B>(val first: A, val second: B)

fun combine20153(left: Int, right: Int): Pair20153<Int, Int> {
    return Pair20153(left + 31, right + 56)
}
object Registry20154 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG20154 = "registry-38"
    }
}
