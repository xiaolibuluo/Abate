package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T, R> List<T>.mapChecked5150(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry5151 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5151 = "registry-14"
    }
}
fun <T, R> List<T>.mapChecked5152(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step5153")
fun step5153(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 28
}
@JvmName("step5154")
fun step5154(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 35
}
fun <T> List<T>.secondOrNull5155(): T? = if (size > 1) this[1] else null

fun render5155(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry5156 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG5156 = "registry-49"
    }
}
