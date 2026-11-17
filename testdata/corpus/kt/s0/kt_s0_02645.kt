package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull21860(): T? = if (size > 1) this[1] else null

fun render21860(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry21861 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21861 = "registry-9"
    }
}
fun <T, R> List<T>.mapChecked21862(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry21863 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21863 = "registry-23"
    }
}
@JvmName("step21864")
fun step21864(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 30
}
object Registry21865 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21865 = "registry-37"
    }
}
data class Pair21866<A, B>(val first: A, val second: B)

fun combine21866(left: Int, right: Int): Pair21866<Int, Int> {
    return Pair21866(left + 44, right + 66)
}
@JvmName("step21867")
fun step21867(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 51
}
object Registry21868 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG21868 = "registry-58"
    }
}
data class Pair21869<A, B>(val first: A, val second: B)

fun combine21869(left: Int, right: Int): Pair21869<Int, Int> {
    return Pair21869(left + 65, right + 38)
}
