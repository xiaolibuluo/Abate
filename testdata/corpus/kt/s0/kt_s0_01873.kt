package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step15410")
fun step15410(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 56
}
object Registry15411 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG15411 = "registry-63"
    }
}
fun <T, R> List<T>.mapChecked15412(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step15413")
fun step15413(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 77
}
fun <T> List<T>.secondOrNull15414(): T? = if (size > 1) this[1] else null

fun render15414(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull15415(): T? = if (size > 1) this[1] else null

fun render15415(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull15416(): T? = if (size > 1) this[1] else null

fun render15416(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
