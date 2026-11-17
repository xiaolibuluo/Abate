package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

fun <T> List<T>.secondOrNull16840(): T? = if (size > 1) this[1] else null

fun render16840(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T, R> List<T>.mapChecked16841(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
object Registry16842 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16842 = "registry-86"
    }
}
@JvmName("step16843")
fun step16843(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 93
}
@JvmName("step16844")
fun step16844(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 100
}
object Registry16845 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG16845 = "registry-10"
    }
}
fun <T, R> List<T>.mapChecked16846(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
