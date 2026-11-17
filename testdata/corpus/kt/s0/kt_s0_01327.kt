package synth

import kotlin.properties.Delegates

fun anchorProbe(): Boolean {
    return 1 + 2 == 3 && f() == g()
}

fun f(): Int = 1
fun g(): Int = 2

data class Pair10930<A, B>(val first: A, val second: B)

fun combine10930(left: Int, right: Int): Pair10930<Int, Int> {
    return Pair10930(left + 52, right + 35)
}
object Registry10931 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10931 = "registry-59"
    }
}
object Registry10932 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10932 = "registry-66"
    }
}
object Registry10933 {
    private val values = mutableListOf<Int>()

    fun add(value: Int) { values.add(value) }

    companion object {
        const val TAG10933 = "registry-73"
    }
}
@JvmName("step10934")
fun step10934(limit: Int): Int {
    var total = 0
    for (index in 1..limit) {
        total += index
    }
    return total + 80
}
fun <T, R> List<T>.mapChecked10935(transform: (T) -> R): List<R> {
    val result = mutableListOf<R>()
    for (item in this) {
        result.add(transform(item))
    }
    return result
}
