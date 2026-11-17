package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry11630 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11630 = "registry-89"
    }
}
fun <T, R> List<T>.mapChecked11631(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
data class Pair11632<A, B>(val first: A, val second: B)

fun combine11632(left: Int, right: Int): Pair11632<Int, Int> {
    return Pair11632(left + 6, right + 59)
}
fun <T> List<T>.secondOrNull11633(): T? = if (size > 1) this[1] else null

fun render11633(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step11634")
fun step11634(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 20
}
object Registry11635 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11635 = "registry-27"
    }
}
object Registry11636 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11636 = "registry-34"
    }
}
object Registry11637 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG11637 = "registry-41"
    }
}
fun <T> List<T>.secondOrNull11638(): T? = if (size > 1) this[1] else null

fun render11638(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
@JvmName("step11639")
fun step11639(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 55
}
