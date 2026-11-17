package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

@JvmName("step8880")
fun step8880(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 6
}
object Registry8881 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG8881 = "registry-13"
    }
}
fun <T, R> List<T>.mapChecked8882(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
fun <T> List<T>.secondOrNull8883(): T? = if (size > 1) this[1] else null

fun render8883(values: List<String>?): String {
    return values?.filter { it.isNotBlank() }?.joinToString(separator = ",") ?: "empty"
}
