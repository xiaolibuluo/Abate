package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

object Registry25750 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25750 = "registry-15"
    }
}
@JvmName("step25751")
fun step25751(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 22
}
fun <T> List<T>.secondOrNull25752(): T? = if (size > 1) this[1] else null

fun render25752(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull25753(): T? = if (size > 1) this[1] else null

fun render25753(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
fun <T> List<T>.secondOrNull25754(): T? = if (size > 1) this[1] else null

fun render25754(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
object Registry25755 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG25755 = "registry-50"
    }
}
fun <T, R> List<T>.mapChecked25756(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T, R> List<T>.mapChecked25757(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
@JvmName("step25758")
fun step25758(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 71
}
fun <T, R> List<T>.mapChecked25759(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
