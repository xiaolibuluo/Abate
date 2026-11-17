package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry3260 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3260 = "registry-72"
    }
}
fun <T> List<T>.secondOrNull3261(): T? = if (size > 1) this[1] else null

fun render3261(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
data class Pair3262<A, B>(val first: A, val second: B)

fun combine3262(left: Int, right: Int): Pair3262<Int, Int> {
    return Pair3262(left + 86, right + 56)
}
fun <T> List<T>.secondOrNull3263(): T? = if (size > 1) this[1] else null

fun render3263(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry3264 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3264 = "registry-100"
    }
}
object Registry3265 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG3265 = "registry-10"
    }
}
fun <T> List<T>.secondOrNull3266(): T? = if (size > 1) this[1] else null

fun render3266(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
